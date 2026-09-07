
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPlannedOrderComponent } from './index.component';
import { PlannedOrderService } from '../../../services/PlannedOrder.service';

describe('IndexPlannedOrderComponent', () => {
  let component: IndexPlannedOrderComponent;
  let fixture: ComponentFixture<IndexPlannedOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPlannedOrderComponent
      ],
      providers: [
        PlannedOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPlannedOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});