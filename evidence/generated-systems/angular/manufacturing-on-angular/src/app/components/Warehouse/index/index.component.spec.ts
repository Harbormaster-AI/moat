
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexWarehouseComponent } from './index.component';
import { WarehouseService } from '../../../services/Warehouse.service';

describe('IndexWarehouseComponent', () => {
  let component: IndexWarehouseComponent;
  let fixture: ComponentFixture<IndexWarehouseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexWarehouseComponent
      ],
      providers: [
        WarehouseService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexWarehouseComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});