
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexGoodsReceiptLineComponent } from './index.component';
import { GoodsReceiptLineService } from '../../../services/GoodsReceiptLine.service';

describe('IndexGoodsReceiptLineComponent', () => {
  let component: IndexGoodsReceiptLineComponent;
  let fixture: ComponentFixture<IndexGoodsReceiptLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexGoodsReceiptLineComponent
      ],
      providers: [
        GoodsReceiptLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexGoodsReceiptLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});