import {RouterModule, Routes} from '@angular/router';
import {AppComponent} from './app.component';
import {ClothesService} from './clothes.service';
import {HomePageComponent} from './home-page/home-page.component';
import {ClothesComponent} from './clothes/clothes.component';
import {ClothFormComponent} from './cloth-form/cloth-form.component';

export const appRoutes: Routes = [
  {path: '', component: HomePageComponent},
  {path: 'clothes', component: ClothesComponent},
  {path: 'clothes/new', component: ClothFormComponent},
  {path: 'clothes/edit/:id', component: ClothFormComponent}
  ];

//https://angular-2-training-book.rangle.io/handout/routing/routeparams.html

// export const AppRouting = RouterModule.forChild(appRoutes);
