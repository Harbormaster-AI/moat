
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLeaveRequestComponent } from './create.component';
import { LeaveRequestService } from '../../../services/LeaveRequest.service';
import { Router } from '@angular/router';

describe('CreateLeaveRequestComponent', () => {
  let component: CreateLeaveRequestComponent;
  let fixture: ComponentFixture<CreateLeaveRequestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLeaveRequestComponent
      ],
      providers: [
        LeaveRequestService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLeaveRequestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});