
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAuditEngagementComponent } from './create.component';
import { AuditEngagementService } from '../../../services/AuditEngagement.service';
import { Router } from '@angular/router';

describe('CreateAuditEngagementComponent', () => {
  let component: CreateAuditEngagementComponent;
  let fixture: ComponentFixture<CreateAuditEngagementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAuditEngagementComponent
      ],
      providers: [
        AuditEngagementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAuditEngagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});