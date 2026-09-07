
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInvestmentAccountComponent } from './create.component';
import { InvestmentAccountService } from '../../../services/InvestmentAccount.service';
import { Router } from '@angular/router';

describe('CreateInvestmentAccountComponent', () => {
  let component: CreateInvestmentAccountComponent;
  let fixture: ComponentFixture<CreateInvestmentAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInvestmentAccountComponent
      ],
      providers: [
        InvestmentAccountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInvestmentAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});