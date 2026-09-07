
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePayoutComponent } from './create.component';
import { PayoutService } from '../../../services/Payout.service';
import { Router } from '@angular/router';

describe('CreatePayoutComponent', () => {
  let component: CreatePayoutComponent;
  let fixture: ComponentFixture<CreatePayoutComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePayoutComponent
      ],
      providers: [
        PayoutService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePayoutComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});