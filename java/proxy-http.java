package example;

import org.apache.http.HttpHost;
import org.apache.http.client.fluent.*;

public class Example {
    public static void main(String[] args) throws Exception {
        HttpHost entry = new HttpHost("gw-us.nstproxy.com", 24125);
        String query = Executor.newInstance()
            .auth(entry, "username", "password")
            .execute(Request.Get("API_URL").viaProxy(entry))
            .returnContent().asString();
        System.out.println(query);
    }
}