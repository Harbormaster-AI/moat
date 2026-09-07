
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTradeComponent } from './create.component';
import { TradeService } from '../../../services/Trade.service';
import { Router } from '@angular/router';

describe('CreateTradeComponent', () => {
  let component: CreateTradeComponent;
  let fixture: ComponentFixture<CreateTradeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTradeComponent
      ],
      providers: [
        TradeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTradeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});