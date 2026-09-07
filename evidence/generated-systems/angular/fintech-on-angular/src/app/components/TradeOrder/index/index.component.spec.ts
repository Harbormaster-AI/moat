
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTradeOrderComponent } from './index.component';
import { TradeOrderService } from '../../../services/TradeOrder.service';

describe('IndexTradeOrderComponent', () => {
  let component: IndexTradeOrderComponent;
  let fixture: ComponentFixture<IndexTradeOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTradeOrderComponent
      ],
      providers: [
        TradeOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTradeOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});