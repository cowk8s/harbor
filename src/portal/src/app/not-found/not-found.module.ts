import { NgModule } from '@angular/core';
import { RouterModule, Routes } from "@angular/router";
import { PageNotFoundComponent } from "./not-found.component";
const routes: Routes = [
  {
    path: '',
    component: PageNotFoundComponent
  }
];
@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  declarations: [PageNotFoundComponent],
})
export class NotFoundModule { }