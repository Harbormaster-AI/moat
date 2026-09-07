
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexOrderItemComponent } from './index.component';
import { OrderItemService } from '../../../services/OrderItem.service';

describe('IndexOrderItemComponent', () => {
  let component: IndexOrderItemComponent;
  let fixture: ComponentFixture<IndexOrderItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexOrderItemComponent
      ],
      providers: [
        OrderItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexOrderItemComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});