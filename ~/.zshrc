# Enhanced sandbox function with multiple options
sandbox() {
    local sandbox_dir="/Users/orocha/Developer/sandbox"
    local current_dir=$(pwd)
    
    # Check if sandbox directory exists
    if [[ ! -d "$sandbox_dir" ]]; then
        echo "Error: Sandbox directory not found at $sandbox_dir"
        return 1
    fi
    
    case "${1:-enter}" in
        "enter"|"")
            echo "🚀 Entering sandbox environment..."
            cd "$sandbox_dir" && make enter-sandbox
            ;;
        "up"|"start")
            echo "🔄 Starting sandbox in background..."
            cd "$sandbox_dir" && make sandbox-up
            ;;
        "down"|"stop")
            echo "🛑 Stopping sandbox..."
            cd "$sandbox_dir" && make sandbox-down
            ;;
        "build"|"rebuild")
            echo "🔨 Building sandbox image..."
            cd "$sandbox_dir" && make sandbox-build
            ;;
        "status")
            echo "📊 Sandbox status:"
            cd "$sandbox_dir" && docker-compose ps
            ;;
        "volumes")
            echo "💾 Sandbox volumes:"
            docker volume ls --filter label=com.docker.compose.project=kanban-app
            ;;
        "help"|"-h"|"--help")
            echo "🛠️  Sandbox Commands:"
            echo "  sandbox [enter]    - Enter interactive sandbox (default)"
            echo "  sandbox up/start   - Start sandbox in background"
            echo "  sandbox down/stop  - Stop sandbox"
            echo "  sandbox build      - Rebuild sandbox image"
            echo "  sandbox status     - Show container status"
            echo "  sandbox volumes    - Show persistent volumes"
            echo "  sandbox help       - Show this help"
            ;;
        *)
            echo "❌ Unknown command: $1"
            echo "Use 'sandbox help' to see available commands"
            return 1
            ;;
    esac
    
    # Return to original directory if command failed
    local exit_code=$?
    if [[ $exit_code -ne 0 ]]; then
        cd "$current_dir"
    fi
    return $exit_code
} 