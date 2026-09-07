
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSupplierComponent } from './index.component';
import { SupplierService } from '../../../services/Supplier.service';

describe('IndexSupplierComponent', () => {
  let component: IndexSupplierComponent;
  let fixture: ComponentFixture<IndexSupplierComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSupplierComponent
      ],
      providers: [
        SupplierService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSupplierComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});