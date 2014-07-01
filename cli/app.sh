ps x | grep animapi | awk '{print $1}' | xargs kill -9
nohup revel run github.com/otiai10/animapi/www prod &
