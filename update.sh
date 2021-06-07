while inotifywait -r -e modify,create,delete,move *; do
    rsync -avz * klaus:battlesnake/
done
