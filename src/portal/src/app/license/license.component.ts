import { Component, OnInit } from '@angular/core';
@Component({
  selector: 'app-license',
  templateUrl: './license.component.html',
})
export class LicenseComponent implements OnInit {
  
  constructor() {}
  public licenseContent: any;
  ngOnInit() {
    this.licenseContent = 'hi';
  }
}