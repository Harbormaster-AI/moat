
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBillingAccountComponent } from './create.component';
import { BillingAccountService } from '../../../services/BillingAccount.service';
import { Router } from '@angular/router';

describe('CreateBillingAccountComponent', () => {
  let component: CreateBillingAccountComponent;
  let fixture: ComponentFixture<CreateBillingAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBillingAccountComponent
      ],
      providers: [
        BillingAccountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBillingAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});