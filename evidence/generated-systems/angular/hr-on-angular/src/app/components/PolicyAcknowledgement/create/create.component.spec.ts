
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePolicyAcknowledgementComponent } from './create.component';
import { PolicyAcknowledgementService } from '../../../services/PolicyAcknowledgement.service';
import { Router } from '@angular/router';

describe('CreatePolicyAcknowledgementComponent', () => {
  let component: CreatePolicyAcknowledgementComponent;
  let fixture: ComponentFixture<CreatePolicyAcknowledgementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePolicyAcknowledgementComponent
      ],
      providers: [
        PolicyAcknowledgementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePolicyAcknowledgementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});