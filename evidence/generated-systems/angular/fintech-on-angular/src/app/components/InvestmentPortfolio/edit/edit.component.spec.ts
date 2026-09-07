
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditInvestmentPortfolioComponent } from './edit.component';
import { InvestmentPortfolioService } from '../../../services/InvestmentPortfolio.service';

describe('EditInvestmentPortfolioComponent', () => {
  let component: EditInvestmentPortfolioComponent;
  let fixture: ComponentFixture<EditInvestmentPortfolioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditInvestmentPortfolioComponent
      ],
      providers: [
        InvestmentPortfolioService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditInvestmentPortfolioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});