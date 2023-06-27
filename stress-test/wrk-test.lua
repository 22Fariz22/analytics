-- init random
math.randomseed(os.time())
request = function()
   url_path = "/analitycs"
   wrk.method = "POST"
  wrk.body   = 
 '{"module" : "settings", "type" : "alert", "event" : "click", "name" : "подтверждение выхода", "data" : {"action" : "cancel"}}' 
  wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
   return wrk.format("POST", url_path,body)
end
