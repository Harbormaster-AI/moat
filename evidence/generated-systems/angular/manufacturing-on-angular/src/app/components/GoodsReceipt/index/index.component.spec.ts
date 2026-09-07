
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexGoodsReceiptComponent } from './index.component';
import { GoodsReceiptService } from '../../../services/GoodsReceipt.service';

describe('IndexGoodsReceiptComponent', () => {
  let component: IndexGoodsReceiptComponent;
  let fixture: ComponentFixture<IndexGoodsReceiptComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexGoodsReceiptComponent
      ],
      providers: [
        GoodsReceiptService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexGoodsReceiptComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});