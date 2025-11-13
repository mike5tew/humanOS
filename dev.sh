#!/bin/bash
# filepath: /Users/michaelstewart/Coding/humanOS/dev.sh
# Quick start development environment

set -e

echo "🚀 Starting HumanOS Development Environment"
echo ""

# Check Docker
if ! command -v docker &> /dev/null; then
    echo "❌ Docker not found. Install Docker to continue."
    exit 1
fi

# Check Node
if ! command -v node &> /dev/null; then
    echo "❌ Node.js not found. Install Node.js v20+ to continue."
    exit 1
fi

echo "✅ Prerequisites check passed"
echo ""

# Start backend
echo "🐳 Starting backend container..."
docker-compose up backend &
BACKEND_PID=$!

# Wait for backend to be ready
echo "⏳ Waiting for backend to be healthy..."
sleep 5

# Check health
if ! curl -s http://localhost:8080/api/health > /dev/null; then
    echo "❌ Backend failed to start"
    kill $BACKEND_PID
    exit 1
fi

echo "✅ Backend is healthy"
echo ""

# Start frontend
echo "🎨 Starting frontend dev server..."
cd frontend
npm run dev &
FRONTEND_PID=$!

echo ""
echo "╔════════════════════════════════════════════════════════╗"
echo "║ HumanOS Development Environment Ready! 🎉              ║"
echo "╠════════════════════════════════════════════════════════╣"
echo "║ Frontend:  http://localhost:5173                       ║"
echo "║ Backend:   http://localhost:8080                       ║"
echo "║ API Docs:  http://localhost:8080/api/health            ║"
echo "╠════════════════════════════════════════════════════════╣"
echo "║ Press Ctrl+C to stop all services                      ║"
echo "╚════════════════════════════════════════════════════════╝"
echo ""

# Wait for interrupt
wait