
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTrainingCourseComponent } from './create.component';
import { TrainingCourseService } from '../../../services/TrainingCourse.service';
import { Router } from '@angular/router';

describe('CreateTrainingCourseComponent', () => {
  let component: CreateTrainingCourseComponent;
  let fixture: ComponentFixture<CreateTrainingCourseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTrainingCourseComponent
      ],
      providers: [
        TrainingCourseService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTrainingCourseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});