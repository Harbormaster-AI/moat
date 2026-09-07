
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTrainingEnrollmentComponent } from './index.component';
import { TrainingEnrollmentService } from '../../../services/TrainingEnrollment.service';

describe('IndexTrainingEnrollmentComponent', () => {
  let component: IndexTrainingEnrollmentComponent;
  let fixture: ComponentFixture<IndexTrainingEnrollmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTrainingEnrollmentComponent
      ],
      providers: [
        TrainingEnrollmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTrainingEnrollmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});