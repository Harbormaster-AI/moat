
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataCategoryComponent } from './create.component';
import { DataCategoryService } from '../../../services/DataCategory.service';
import { Router } from '@angular/router';

describe('CreateDataCategoryComponent', () => {
  let component: CreateDataCategoryComponent;
  let fixture: ComponentFixture<CreateDataCategoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataCategoryComponent
      ],
      providers: [
        DataCategoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataCategoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});