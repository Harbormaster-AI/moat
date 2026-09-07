
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRoleAssignmentComponent } from './index.component';
import { RoleAssignmentService } from '../../../services/RoleAssignment.service';

describe('IndexRoleAssignmentComponent', () => {
  let component: IndexRoleAssignmentComponent;
  let fixture: ComponentFixture<IndexRoleAssignmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRoleAssignmentComponent
      ],
      providers: [
        RoleAssignmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRoleAssignmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});