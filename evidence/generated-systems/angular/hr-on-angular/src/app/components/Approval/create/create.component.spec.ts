
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateApprovalComponent } from './create.component';
import { ApprovalService } from '../../../services/Approval.service';
import { Router } from '@angular/router';

describe('CreateApprovalComponent', () => {
  let component: CreateApprovalComponent;
  let fixture: ComponentFixture<CreateApprovalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateApprovalComponent
      ],
      providers: [
        ApprovalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateApprovalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});