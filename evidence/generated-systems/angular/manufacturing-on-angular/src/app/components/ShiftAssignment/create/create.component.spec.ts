
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateShiftAssignmentComponent } from './create.component';
import { ShiftAssignmentService } from '../../../services/ShiftAssignment.service';
import { Router } from '@angular/router';

describe('CreateShiftAssignmentComponent', () => {
  let component: CreateShiftAssignmentComponent;
  let fixture: ComponentFixture<CreateShiftAssignmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateShiftAssignmentComponent
      ],
      providers: [
        ShiftAssignmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateShiftAssignmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});