from locust import HttpUser, task

class ProductService(HttpUser):
    @task
    def service(self):
        self.client.get("/")
        self.client.get("/")

class OrderService (HttpUser):
    @task
    def service(self):
        self.client.get("/")

class LoginService(HttpUser):
    @task
    def service(self):
        self.client.get("/")