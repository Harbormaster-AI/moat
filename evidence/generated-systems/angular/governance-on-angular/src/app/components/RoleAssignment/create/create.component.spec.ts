
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRoleAssignmentComponent } from './create.component';
import { RoleAssignmentService } from '../../../services/RoleAssignment.service';
import { Router } from '@angular/router';

describe('CreateRoleAssignmentComponent', () => {
  let component: CreateRoleAssignmentComponent;
  let fixture: ComponentFixture<CreateRoleAssignmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRoleAssignmentComponent
      ],
      providers: [
        RoleAssignmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRoleAssignmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});