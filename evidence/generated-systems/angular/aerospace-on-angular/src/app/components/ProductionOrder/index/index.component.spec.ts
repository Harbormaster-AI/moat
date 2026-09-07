
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexProductionOrderComponent } from './index.component';
import { ProductionOrderService } from '../../../services/ProductionOrder.service';

describe('IndexProductionOrderComponent', () => {
  let component: IndexProductionOrderComponent;
  let fixture: ComponentFixture<IndexProductionOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexProductionOrderComponent
      ],
      providers: [
        ProductionOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexProductionOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});