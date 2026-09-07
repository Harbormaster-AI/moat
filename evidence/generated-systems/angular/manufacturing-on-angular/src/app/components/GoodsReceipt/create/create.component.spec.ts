
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateGoodsReceiptComponent } from './create.component';
import { GoodsReceiptService } from '../../../services/GoodsReceipt.service';
import { Router } from '@angular/router';

describe('CreateGoodsReceiptComponent', () => {
  let component: CreateGoodsReceiptComponent;
  let fixture: ComponentFixture<CreateGoodsReceiptComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateGoodsReceiptComponent
      ],
      providers: [
        GoodsReceiptService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateGoodsReceiptComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});