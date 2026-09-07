
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInventoryItemComponent } from './index.component';
import { InventoryItemService } from '../../../services/InventoryItem.service';

describe('IndexInventoryItemComponent', () => {
  let component: IndexInventoryItemComponent;
  let fixture: ComponentFixture<IndexInventoryItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInventoryItemComponent
      ],
      providers: [
        InventoryItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInventoryItemComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});