
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEmploymentAssignmentComponent } from './index.component';
import { EmploymentAssignmentService } from '../../../services/EmploymentAssignment.service';

describe('IndexEmploymentAssignmentComponent', () => {
  let component: IndexEmploymentAssignmentComponent;
  let fixture: ComponentFixture<IndexEmploymentAssignmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEmploymentAssignmentComponent
      ],
      providers: [
        EmploymentAssignmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEmploymentAssignmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});