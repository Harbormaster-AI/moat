
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSalesOrderComponent } from './index.component';
import { SalesOrderService } from '../../../services/SalesOrder.service';

describe('IndexSalesOrderComponent', () => {
  let component: IndexSalesOrderComponent;
  let fixture: ComponentFixture<IndexSalesOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSalesOrderComponent
      ],
      providers: [
        SalesOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSalesOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});