import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    component: {},
    children: [
      { path: '', redirectTo: 'projects', pathMatch: 'full' },
      {
        path: 'project',
        loadChildren: () => import('./')
      }
    ]
  }
]
@NgModule({
  imports: [
    RouterModule.forChild(routes),
  ],
  declarations: [
    
  ]
})
export class BaseModule {

}