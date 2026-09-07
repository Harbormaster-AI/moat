
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDependentComponent } from './create.component';
import { DependentService } from '../../../services/Dependent.service';
import { Router } from '@angular/router';

describe('CreateDependentComponent', () => {
  let component: CreateDependentComponent;
  let fixture: ComponentFixture<CreateDependentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDependentComponent
      ],
      providers: [
        DependentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDependentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});