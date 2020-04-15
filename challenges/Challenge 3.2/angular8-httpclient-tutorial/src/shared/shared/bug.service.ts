import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from '../bug';
import { Observable, throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BugService {
  //baseurl = 'http://localhost:3000';
  baseurl='/api/webapi/v1/user'
  constructor(private http: HttpClient) { }

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json','Access-Control-Allow-Origin': '*'

    })
}
    /*
httpOptions={
  headers:new HttpHeaders()
  .set('Content-Type','application/json')
  .set('Cache-Control','no-cache')
  .set('Access-Control-Allow-Origin', '*')
  .set('method', 'POST')
}*/
    // POST
    CreateBug(data): Observable<User> {
      return this.http.post<User>(this.baseurl, JSON.stringify(data), this.httpOptions )
      .pipe(
        retry(1),
        catchError(this.errorHandl)
      )
    }  
  
    // GET
    GetIssue(id): Observable<User> {
      return this.http.get<User>(this.baseurl+'/'+ id)
      .pipe(
       retry(1),
        catchError(this.errorHandl)
      );
    }
  
    // GET
    GetIssues(): Observable<User[]> {
      return this.http.get<User[]>(this.baseurl)
      .pipe(
        retry(1),
         catchError(this.errorHandl)
       );
     }
  
    // PUT
    UpdateBug(id, data): Observable<User> {
      return this.http.put<User>(this.baseurl + '/' + id, JSON.stringify(data), this.httpOptions)
      .pipe(
        retry(1),
        catchError(this.errorHandl)
      )
    }
  
    // DELETE
    DeleteBug(id){
      return this.http.delete<User>(this.baseurl + '/' + id, this.httpOptions)
      .pipe(
        retry(1),
        catchError(this.errorHandl)
      )
    }
  
    // Error handling
    errorHandl(error) {
       let errorMessage = '';
       if(error.error instanceof ErrorEvent) {
         // Get client-side error
         errorMessage = error.error.message;
       } else {
         // Get server-side error
         errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
       }
       console.log(errorMessage);
       return throwError(errorMessage);
    }
  
}
