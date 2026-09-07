
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEmploymentAssignmentComponent } from './create.component';
import { EmploymentAssignmentService } from '../../../services/EmploymentAssignment.service';
import { Router } from '@angular/router';

describe('CreateEmploymentAssignmentComponent', () => {
  let component: CreateEmploymentAssignmentComponent;
  let fixture: ComponentFixture<CreateEmploymentAssignmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEmploymentAssignmentComponent
      ],
      providers: [
        EmploymentAssignmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEmploymentAssignmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});