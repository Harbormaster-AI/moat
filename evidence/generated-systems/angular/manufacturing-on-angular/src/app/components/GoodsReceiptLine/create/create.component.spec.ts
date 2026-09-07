
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateGoodsReceiptLineComponent } from './create.component';
import { GoodsReceiptLineService } from '../../../services/GoodsReceiptLine.service';
import { Router } from '@angular/router';

describe('CreateGoodsReceiptLineComponent', () => {
  let component: CreateGoodsReceiptLineComponent;
  let fixture: ComponentFixture<CreateGoodsReceiptLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateGoodsReceiptLineComponent
      ],
      providers: [
        GoodsReceiptLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateGoodsReceiptLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});