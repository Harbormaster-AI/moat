
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInventoryTransactionComponent } from './index.component';
import { InventoryTransactionService } from '../../../services/InventoryTransaction.service';

describe('IndexInventoryTransactionComponent', () => {
  let component: IndexInventoryTransactionComponent;
  let fixture: ComponentFixture<IndexInventoryTransactionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInventoryTransactionComponent
      ],
      providers: [
        InventoryTransactionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInventoryTransactionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});