
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDirectDebitMandateComponent } from './create.component';
import { DirectDebitMandateService } from '../../../services/DirectDebitMandate.service';
import { Router } from '@angular/router';

describe('CreateDirectDebitMandateComponent', () => {
  let component: CreateDirectDebitMandateComponent;
  let fixture: ComponentFixture<CreateDirectDebitMandateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDirectDebitMandateComponent
      ],
      providers: [
        DirectDebitMandateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDirectDebitMandateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});