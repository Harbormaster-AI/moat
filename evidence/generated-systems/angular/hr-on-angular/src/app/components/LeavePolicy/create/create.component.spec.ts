
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLeavePolicyComponent } from './create.component';
import { LeavePolicyService } from '../../../services/LeavePolicy.service';
import { Router } from '@angular/router';

describe('CreateLeavePolicyComponent', () => {
  let component: CreateLeavePolicyComponent;
  let fixture: ComponentFixture<CreateLeavePolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLeavePolicyComponent
      ],
      providers: [
        LeavePolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLeavePolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});