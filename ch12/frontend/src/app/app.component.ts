import { Component, Inject, ViewChild, OnInit } from '@angular/core';
import { NgbModal, NgbModalRef } from '@ng-bootstrap/ng-bootstrap';
import { CognitoService } from './services/cognito.service';
import {LOCAL_STORAGE, WebStorageService} from 'angular-webstorage-service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  private loginModal: NgbModalRef;
  public loginError : boolean;

  @ViewChild('contentlogin') 
  private content;

  constructor(private modalService: NgbModal,
              private cognitoService: CognitoService,
              @Inject(LOCAL_STORAGE) private storage: WebStorageService){}

  ngOnInit(){
    if(!this.storage.get("COGNITO_TOKEN")){
      this.loginModal = this.modalService.open(this.content)
    }
  }

  signin(username, password){
    this.cognitoService.auth(username, password, (err, token) => {
      if(err){
        this.loginError = true
      }else{
        this.loginError = false
        this.storage.set("COGNITO_TOKEN", token)
        this.loginModal.close()
        window.location.reload()
      }
    })
  }
}
