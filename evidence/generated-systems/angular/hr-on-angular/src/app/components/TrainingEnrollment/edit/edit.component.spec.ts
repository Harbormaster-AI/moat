
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditTrainingEnrollmentComponent } from './edit.component';
import { TrainingEnrollmentService } from '../../../services/TrainingEnrollment.service';

describe('EditTrainingEnrollmentComponent', () => {
  let component: EditTrainingEnrollmentComponent;
  let fixture: ComponentFixture<EditTrainingEnrollmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditTrainingEnrollmentComponent
      ],
      providers: [
        TrainingEnrollmentService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditTrainingEnrollmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});