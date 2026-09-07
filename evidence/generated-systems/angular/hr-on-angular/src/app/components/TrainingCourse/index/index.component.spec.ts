
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTrainingCourseComponent } from './index.component';
import { TrainingCourseService } from '../../../services/TrainingCourse.service';

describe('IndexTrainingCourseComponent', () => {
  let component: IndexTrainingCourseComponent;
  let fixture: ComponentFixture<IndexTrainingCourseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTrainingCourseComponent
      ],
      providers: [
        TrainingCourseService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTrainingCourseComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});