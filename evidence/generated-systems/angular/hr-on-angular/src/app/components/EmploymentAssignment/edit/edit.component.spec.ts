
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditEmploymentAssignmentComponent } from './edit.component';
import { EmploymentAssignmentService } from '../../../services/EmploymentAssignment.service';

describe('EditEmploymentAssignmentComponent', () => {
  let component: EditEmploymentAssignmentComponent;
  let fixture: ComponentFixture<EditEmploymentAssignmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditEmploymentAssignmentComponent
      ],
      providers: [
        EmploymentAssignmentService,
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

    fixture = TestBed.createComponent(EditEmploymentAssignmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});