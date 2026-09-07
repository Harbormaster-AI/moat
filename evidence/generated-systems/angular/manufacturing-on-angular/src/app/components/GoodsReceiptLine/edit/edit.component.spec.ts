
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditGoodsReceiptLineComponent } from './edit.component';
import { GoodsReceiptLineService } from '../../../services/GoodsReceiptLine.service';

describe('EditGoodsReceiptLineComponent', () => {
  let component: EditGoodsReceiptLineComponent;
  let fixture: ComponentFixture<EditGoodsReceiptLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditGoodsReceiptLineComponent
      ],
      providers: [
        GoodsReceiptLineService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditGoodsReceiptLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});