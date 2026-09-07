
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSalesOrderLineComponent } from './index.component';
import { SalesOrderLineService } from '../../../services/SalesOrderLine.service';

describe('IndexSalesOrderLineComponent', () => {
  let component: IndexSalesOrderLineComponent;
  let fixture: ComponentFixture<IndexSalesOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSalesOrderLineComponent
      ],
      providers: [
        SalesOrderLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSalesOrderLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});