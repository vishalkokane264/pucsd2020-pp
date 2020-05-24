import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-after-sign-in',
  templateUrl: './after-sign-in.component.html',
  styleUrls: ['./after-sign-in.component.css']
})
export class AfterSignInComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
