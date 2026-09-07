
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCarePlanComponent } from './create.component';
import { CarePlanService } from '../../../services/CarePlan.service';
import { Router } from '@angular/router';

describe('CreateCarePlanComponent', () => {
  let component: CreateCarePlanComponent;
  let fixture: ComponentFixture<CreateCarePlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCarePlanComponent
      ],
      providers: [
        CarePlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCarePlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});