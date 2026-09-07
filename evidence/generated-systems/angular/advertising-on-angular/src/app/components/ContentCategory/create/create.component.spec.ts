
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateContentCategoryComponent } from './create.component';
import { ContentCategoryService } from '../../../services/ContentCategory.service';
import { Router } from '@angular/router';

describe('CreateContentCategoryComponent', () => {
  let component: CreateContentCategoryComponent;
  let fixture: ComponentFixture<CreateContentCategoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateContentCategoryComponent
      ],
      providers: [
        ContentCategoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateContentCategoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});