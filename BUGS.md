- Multiple subscriptions per session to the same query are not cleaned up properly on disconnect. On next update a n-1 subscriptions will persist, but they will all have closed channels and the first send to any of these will result in a crash.