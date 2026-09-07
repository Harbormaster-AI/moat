
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTrainingEnrollmentComponent } from './create.component';
import { TrainingEnrollmentService } from '../../../services/TrainingEnrollment.service';
import { Router } from '@angular/router';

describe('CreateTrainingEnrollmentComponent', () => {
  let component: CreateTrainingEnrollmentComponent;
  let fixture: ComponentFixture<CreateTrainingEnrollmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTrainingEnrollmentComponent
      ],
      providers: [
        TrainingEnrollmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTrainingEnrollmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});