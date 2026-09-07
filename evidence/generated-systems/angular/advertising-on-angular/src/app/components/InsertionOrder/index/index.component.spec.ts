
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInsertionOrderComponent } from './index.component';
import { InsertionOrderService } from '../../../services/InsertionOrder.service';

describe('IndexInsertionOrderComponent', () => {
  let component: IndexInsertionOrderComponent;
  let fixture: ComponentFixture<IndexInsertionOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInsertionOrderComponent
      ],
      providers: [
        InsertionOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInsertionOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});