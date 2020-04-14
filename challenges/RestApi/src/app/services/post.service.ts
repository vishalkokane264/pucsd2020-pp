import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class PostService {
  private url='http://localhost:9090/webapi/v1/user/1';

  constructor(private httpClient:HttpClient) { }
  public getSendRequest(){
    return this.httpClient.get(this.url);
  }
}
