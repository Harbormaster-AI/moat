
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTradeOrderComponent } from './create.component';
import { TradeOrderService } from '../../../services/TradeOrder.service';
import { Router } from '@angular/router';

describe('CreateTradeOrderComponent', () => {
  let component: CreateTradeOrderComponent;
  let fixture: ComponentFixture<CreateTradeOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTradeOrderComponent
      ],
      providers: [
        TradeOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTradeOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});