from locust import HttpUser, task
class OrderService (HttpUser):
    @task
    def service(self):
        self.client.post("/v1/orders/get-order",json=
           {
              "order_id": "cf3e4c95-a1bb-48fa-90b1-72b0ce43afc8"
            })

class LoginService(HttpUser):
    @task
    def service(self):
        self.client.get("/")
